import React from 'react';
import PropTypes from 'prop-types';
import { List, ListItem } from './List';
import ListItemText from '@material-ui/core/ListItemText';
import { useTranslation } from 'react-i18next';

import { formatDayTime, formatWeekdays } from '../schedule';

export default function IntervalOptionsList(props) {
  const {
    weekdays,
    fromDayTime,
    toDayTime,
    onWeekdaysClick,
    onFromDayTimeClick,
    onToDayTimeClick
  } = props;

  const { t } = useTranslation();

  return (
    <List>
      <ListItem onClick={onWeekdaysClick}>
        <ListItemText primary={t('weekdays')} secondary={formatWeekdays(weekdays, t)} />
      </ListItem>
      <ListItem onClick={onFromDayTimeClick}>
        <ListItemText primary={t('start-time')} secondary={formatDayTime(fromDayTime, t)} />
      </ListItem>
      <ListItem onClick={onToDayTimeClick}>
        <ListItemText primary={t('end-time')} secondary={formatDayTime(toDayTime, t)} />
      </ListItem>
    </List>
  );
}

IntervalOptionsList.propTypes = {
    weekdays: PropTypes.array,
    fromDayTime: PropTypes.object,
    toDayTime: PropTypes.object,
    onWeekdaysClick: PropTypes.func.isRequired,
    onFromDayTimeClick: PropTypes.func.isRequired,
    onToDayTimeClick: PropTypes.func.isRequired,
};
