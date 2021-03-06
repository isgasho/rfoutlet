import React from 'react';
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/core/styles';
import { TimePicker } from '@material-ui/pickers';
import { useTranslation } from 'react-i18next';

const useStyles = makeStyles({
  timePicker: {
    position: 'fixed',
    top: -1000,
  },
});

export default function IntervalTimePicker(props) {
  const classes = useStyles();

  const { t } = useTranslation();

  return (
    <TimePicker
      className={classes.timePicker}
      clearable
      open
      ampm={false}
      okLabel={t('picker-label-ok')}
      cancelLabel={t('picker-label-cancel')}
      clearLabel={t('picker-label-clear')}
      {...props}
    />
  );
}

IntervalTimePicker.propTypes = {
  onChange: PropTypes.func.isRequired,
  onClose: PropTypes.func.isRequired,
};
