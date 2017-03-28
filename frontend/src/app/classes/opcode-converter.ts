import { OpcodeCollection } from '../interfaces/opcode-collection';
import opcodeCollection from '../../../../protocol/protocol.json';

export class OpcodeConverter {
  private toOpcode: OpcodeCollection;
  private toIdentifier: OpcodeCollection;

  constructor() {
    this.toOpcode = opcodeCollection;
    this.toIdentifier = {};
    for(let i in this.toOpcode) {
      this.toIdentifier[this.toOpcode[i].id] = {
        id: i,
        methods: {}
      };
      for(let j in this.toOpcode[i].methods) {
        this.toIdentifier[this.toOpcode[i].id].methods[this.toOpcode[i].methods[j]] = j;
      }
    }
  }

  public getOpcodes(pkg: string, message: string): { pkg: string, message: string } {
    return {
      pkg: this.toOpcode[pkg].id,
      message: this.toOpcode[pkg].methods[message]
    };
  }

  public getIdentifiers(pkg: string, message: string): { pkg: string, message: string } {
    return {
      pkg: this.toIdentifier[pkg].id,
      message:this.toIdentifier[pkg].methods[message]
    }
  }
}
