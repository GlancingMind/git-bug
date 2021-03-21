import React from 'react';
import { Link } from 'react-router-dom';

import TableCell from '@material-ui/core/TableCell/TableCell';
import TableRow from '@material-ui/core/TableRow/TableRow';
import Tooltip from '@material-ui/core/Tooltip/Tooltip';
import { makeStyles } from '@material-ui/core/styles';
import CheckCircleOutline from '@material-ui/icons/CheckCircleOutline';
import ErrorOutline from '@material-ui/icons/ErrorOutline';

import Author from 'src/components/Author';
import Date from 'src/components/Date';
import Label from 'src/components/Label';
import { Status } from 'src/gqlTypes';

import { BugRowFragment } from './BugRow.generated';

type OpenClosedProps = { className: string };
const Open = ({ className }: OpenClosedProps) => (
  <Tooltip title="Open">
    <ErrorOutline htmlColor="#28a745" className={className} />
  </Tooltip>
);

const Closed = ({ className }: OpenClosedProps) => (
  <Tooltip title="Closed">
    <CheckCircleOutline htmlColor="#cb2431" className={className} />
  </Tooltip>
);

type StatusProps = { className: string; status: Status };
const BugStatus: React.FC<StatusProps> = ({
  status,
  className,
}: StatusProps) => {
  switch (status) {
    case 'OPEN':
      return <Open className={className} />;
    case 'CLOSED':
      return <Closed className={className} />;
    default:
      return <p>{'unknown status ' + status}</p>;
  }
};

const useStyles = makeStyles((theme) => ({
  cell: {
    display: 'flex',
    alignItems: 'center',
    padding: theme.spacing(1),
    '& a': {
      textDecoration: 'none',
    },
  },
  status: {
    margin: theme.spacing(1, 2),
  },
  expand: {
    width: '100%',
    lineHeight: '20px',
  },
  title: {
    display: 'inline',
    color: theme.palette.text.primary,
    fontSize: '1.3rem',
    fontWeight: 500,
  },
  details: {
    lineHeight: '1.5rem',
    color: theme.palette.text.secondary,
  },
  labels: {
    paddingLeft: theme.spacing(1),
    '& > *': {
      display: 'inline-block',
    },
  },
}));

type Props = {
  bug: BugRowFragment;
};

function BugRow({ bug }: Props) {
  const classes = useStyles();
  return (
    <TableRow hover>
      <TableCell className={classes.cell}>
        <BugStatus status={bug.status} className={classes.status} />
        <div className={classes.expand}>
          <Link to={'bug/' + bug.humanId}>
            <div className={classes.expand}>
              <span className={classes.title}>{bug.title}</span>
              {bug.labels.length > 0 && (
                <span className={classes.labels}>
                  {bug.labels.map((l) => (
                    <Label key={l.name} label={l} />
                  ))}
                </span>
              )}
            </div>
          </Link>
          <div className={classes.details}>
            {bug.humanId} opened&nbsp;
            <Date date={bug.createdAt} />
            &nbsp;by
            <Author className={classes.details} author={bug.author} />
          </div>
        </div>
      </TableCell>
    </TableRow>
  );
}

export default BugRow;
