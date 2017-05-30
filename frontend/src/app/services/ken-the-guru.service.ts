import { Injectable } from '@angular/core';
import { Observable, Subject, Subscription } from 'rxjs/Rx';
import { PartialObserver } from 'rxjs/Observer';
import { SocketService } from './socket.service'
import { ProtoResponse } from '../interfaces/proto-response';

const SOCKET_URL = 'ws://localhost:8083';
const KENTHEGURU_TYPE = 'kentheguru';

@Injectable()
export class KenTheGuruService {
  public messages: Subject<object>;

  constructor(private wsService: SocketService) {
    this.messages = <Subject<ProtoResponse>>this.wsService
      .connect(SOCKET_URL)
      .map((response: MessageEvent): ProtoResponse => {
        return this.wsService.decodeMessage(new Uint8Array(response.data), KENTHEGURU_TYPE);
      });
  }

  public next(message: string, data: object) {
    this.messages.next(this.wsService.encodeMessage(KENTHEGURU_TYPE, message, data));
  }

}
