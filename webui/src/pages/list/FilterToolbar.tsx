import { pipe } from '@arrows/composition';
import { LocationDescriptor } from 'history';
import React from 'react';

import Toolbar from '@material-ui/core/Toolbar';
import { makeStyles } from '@material-ui/core/styles';
import CheckCircleOutline from '@material-ui/icons/CheckCircleOutline';
import ErrorOutline from '@material-ui/icons/ErrorOutline';

import {
  Filter,
  FilterDropdown,
  FilterProps,
  parse,
  Query,
  stringify,
} from './Filter';
import { useBugCountQuery } from './FilterToolbar.generated';
import { useListIdentitiesQuery } from './ListIdentities.generated';
import { useListLabelsQuery } from './ListLabels.generated';

const useStyles = makeStyles((theme) => ({
  toolbar: {
    backgroundColor: theme.palette.primary.light,
    borderColor: theme.palette.divider,
    borderWidth: '1px 0',
    borderStyle: 'solid',
    margin: theme.spacing(0, -1),
  },
  spacer: {
    flex: 1,
  },
}));

// This prepends the filter text with a count
type CountingFilterProps = {
  query: string; // the query used as a source to count the number of element
  children: React.ReactNode;
} & FilterProps;

function CountingFilter({ query, children, ...props }: CountingFilterProps) {
  const { data, loading, error } = useBugCountQuery({
    variables: { query },
  });

  var prefix;
  if (loading) prefix = '...';
  else if (error || !data?.repository) prefix = '???';
  // TODO: better prefixes & error handling
  else prefix = data.repository.bugs.totalCount;

  return (
    <Filter {...props}>
      {prefix} {children}
    </Filter>
  );
}

type Props = {
  query: string;
  queryLocation: (query: string) => LocationDescriptor;
};

function FilterToolbar({ query, queryLocation }: Props) {
  const classes = useStyles();
  const params: Query = parse(query);
  const { data: identitiesData } = useListIdentitiesQuery();
  const { data: labelsData } = useListLabelsQuery();

  let identities: any = [];
  let labels: any = [];

  if (
    identitiesData?.repository &&
    identitiesData.repository.allIdentities &&
    identitiesData.repository.allIdentities.nodes
  ) {
    identities = identitiesData.repository.allIdentities.nodes.map((node) => [
      node.name,
      node.name,
    ]);
  }

  if (
    labelsData?.repository &&
    labelsData.repository.validLabels &&
    labelsData.repository.validLabels.nodes
  ) {
    labels = labelsData.repository.validLabels.nodes.map((node) => [
      node.name,
      node.name,
    ]);
  }

  const hasKey = (key: string): boolean =>
    params[key] && params[key].length > 0;
  const hasValue = (key: string, value: string): boolean =>
    hasKey(key) && params[key].includes(value);
  const containsValue = (key: string, value: string): boolean =>
    hasKey(key) && params[key].indexOf(value) !== -1;
  const loc = pipe(stringify, queryLocation);
  const replaceParam = (key: string, value: string) => (
    params: Query
  ): Query => ({
    ...params,
    [key]: [value],
  });
  const toggleParam = (key: string, value: string) => (
    params: Query
  ): Query => ({
    ...params,
    [key]: params[key] && params[key].includes(value) ? [] : [value],
  });
  const toggleOrAddParam = (key: string, value: string) => (
    params: Query
  ): Query => {
    const values = params[key];
    return {
      ...params,
      [key]:
        params[key] && params[key].includes(value)
          ? values.filter((v) => v !== value)
          : values
          ? [...values, value]
          : [value],
    };
  };
  const clearParam = (key: string) => (params: Query): Query => ({
    ...params,
    [key]: [],
  });

  // TODO: author/label filters
  return (
    <Toolbar className={classes.toolbar}>
      <CountingFilter
        active={hasValue('status', 'open')}
        query={pipe(
          replaceParam('status', 'open'),
          clearParam('sort'),
          stringify
        )(params)}
        to={pipe(toggleParam('status', 'open'), loc)(params)}
        icon={ErrorOutline}
      >
        open
      </CountingFilter>
      <CountingFilter
        active={hasValue('status', 'closed')}
        query={pipe(
          replaceParam('status', 'closed'),
          clearParam('sort'),
          stringify
        )(params)}
        to={pipe(toggleParam('status', 'closed'), loc)(params)}
        icon={CheckCircleOutline}
      >
        closed
      </CountingFilter>
      <div className={classes.spacer} />
      {/*
      <Filter active={hasKey('author')}>Author</Filter>
      <Filter active={hasKey('label')}>Label</Filter>
      */}
      <FilterDropdown
        dropdown={identities}
        itemActive={(key) => hasValue('author', key)}
        to={(key) => pipe(toggleParam('author', key), loc)(params)}
        hasFilter
      >
        Author
      </FilterDropdown>
      <FilterDropdown
        dropdown={labels}
        itemActive={(key) => containsValue('label', key)}
        to={(key) => pipe(toggleOrAddParam('label', key), loc)(params)}
        hasFilter
      >
        Labels
      </FilterDropdown>
      <FilterDropdown
        dropdown={[
          ['id', 'ID'],
          ['creation', 'Newest'],
          ['creation-asc', 'Oldest'],
          ['edit', 'Recently updated'],
          ['edit-asc', 'Least recently updated'],
        ]}
        itemActive={(key) => hasValue('sort', key)}
        to={(key) => pipe(toggleParam('sort', key), loc)(params)}
      >
        Sort
      </FilterDropdown>
    </Toolbar>
  );
}

export default FilterToolbar;
