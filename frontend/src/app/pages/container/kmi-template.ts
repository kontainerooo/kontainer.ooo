import { GlobalDataService } from '../../services/global-data.service';


export class KmiTemplate {
  constructor(private gds: GlobalDataService) {
  }

  getParameter(name: string): string {
    // TODO implement when service is ready
    return 'aParameter';
  }

  sendCommand(name: string) {
    
  }

  getFiles(directory: string) {

  }

  getFile(path: string) {

  }

  deleteFile(path: string) {

  }

  deleteDirectory(directory: string) {

  }

  uploadPublicKey(pubKey: string) {

  }
}
