/* eslint global-require: 0 */

import { fromJS } from 'immutable';

export default fromJS([
  require('../schema/email.json'),
  require('../schema/sms.json'),
]);
