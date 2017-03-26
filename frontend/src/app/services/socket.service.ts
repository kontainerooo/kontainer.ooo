import { Injectable } from '@angular/core';
import * as Rx from 'rxjs/Rx';
import * as pb from '../../messages/messages';
import { OpcodeConverter } from '../classes/opcode-converter';

@Injectable()
export class SocketService {
  private socket: Rx.Subject<MessageEvent>;
  private opcodeConverter: OpcodeConverter;
  private readonly PROTOCOL_VERSION: string;

  constructor() { 
    this.opcodeConverter = new OpcodeConverter();
    this.PROTOCOL_VERSION = 'kroov1';
  }

  public connect(url): Rx.Subject<MessageEvent> {
    if(!this.socket) {
      this.socket = this.create(url);
    }

    return this.socket;
  }

  private create(url): Rx.Subject<MessageEvent> {
    // TODO add protocol when supported
    let ws = new WebSocket(url/*, this.PROTOCOL_VERSION*/);
    ws.binaryType = 'arraybuffer';

    let observable: Rx.Observable<MessageEvent> = Rx.Observable.create(
      (obs: Rx.Observer<MessageEvent>) => {
        ws.onmessage = obs.next.bind(obs);
        ws.onerror = obs.error.bind(obs);
        ws.onclose = obs.complete.bind(obs);

        return ws.close.bind(ws);
      }
    );

    let observer = {
      next: (data: Uint8Array) => {
        if (ws.readyState === WebSocket.OPEN) {
          ws.send(data);
        }
      }
    };

    return Rx.Subject.create(observer, observable);
  }

  private generateOpcode(opcode: string): Array<number> {
    const opcodeLength: number = 3;
    if (opcode.length === opcodeLength) {
      let opcodeArray: Array<number> = new Array<number>();
      for (let i = 0; i < opcodeLength; i++) {
        opcodeArray[i] = opcode.charCodeAt(i);
      }
      return opcodeArray;
    }
    return null;
  }

  public encodeMessage(pkg: string, message: string, data: object): Uint8Array {
    const opcodes: { pkg: string, message: string } = this.opcodeConverter.getOpcodes(pkg, message.replace(/Request|Response/, ''));
    const pkgArray: Array<number> = this.generateOpcode(opcodes.pkg);
    const messageArray: Array<number> = this.generateOpcode(opcodes.message);
    const dataArray: Array<number> = Array.from(<Uint8Array>pb[pkg][message].encode(data).finish());
    if (pkgArray !== null && messageArray !== null) {
      return new Uint8Array([...pkgArray, ...messageArray, ...dataArray]);
    }
  }

  public decodeMessage(encodedMessage: Uint8Array, pkgWanted: string): object {
    console.log(encodedMessage);
    const pkgArray: Uint8Array = encodedMessage.slice(0, 3);
    const messageArray: Uint8Array = encodedMessage.slice(3, 6);
    let pkg: string = '';
    let message: string = '';
    for(let i in pkgArray) {
      pkg += String.fromCharCode(pkgArray[i]);
      message += String.fromCharCode(messageArray[i]);
    }
    console.log(pkg, message);
    const identifiers: { pkg: string, message: string } = this.opcodeConverter.getIdentifiers(pkg, message);
    if(identifiers.pkg == pkgWanted) {
      return pb[identifiers.pkg][`${identifiers.message}Response`].decode(encodedMessage.slice(6)).toObject();
    } else {
      return null;
    }
  }
}