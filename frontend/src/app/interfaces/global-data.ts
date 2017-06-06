import { user, kmi } from '../../messages/messages';

export interface GlobalData {
  user?: user.User,
  KMDI?: kmi.KMDI[]
}
