// TODO example import as there is no service for this right now
import { KmiService } from '../../services/kmi.service';


export class KmiTemplate {
  private cKmiService: KmiService;

  constructor(kmiService: KmiService) {
    this.cKmiService = kmiService;
  }

  getParameter(name: string): string {
    // TODO implement when service is ready
    return 'aParameter';
  }

  executeCommand(name: string) {
    
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
