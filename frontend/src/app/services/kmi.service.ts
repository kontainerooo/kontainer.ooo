import { Injectable } from '@angular/core';
import { Observable, Subject, Subscription } from 'rxjs/Rx';
import { PartialObserver } from 'rxjs/Observer';
import { SocketService } from './socket.service'
import { ProtoResponse } from '../interfaces/proto-response';

const SOCKET_URL = 'ws://localhost:8083';
const USER_TYPE = 'kmi';

@Injectable()
export class KmiService {
  public messages: Subject<object>;

  constructor(private wsService: SocketService) {
    this.messages = <Subject<ProtoResponse>>this.wsService
      .connect(SOCKET_URL)
      .map((response: MessageEvent): ProtoResponse => {
        return this.wsService.decodeMessage(new Uint8Array(response.data), USER_TYPE);
      });
  }

  public reconnect(): Subject<object> {
    this.messages = <Subject<Uint8Array>>this.wsService
      .connect(SOCKET_URL)
      .map((response: MessageEvent): object => {
        return this.wsService.decodeMessage(new Uint8Array(response.data), USER_TYPE);
      });
    return this.messages;
  }

  public next(message: string, data: object) {
    this.messages.next(this.wsService.encodeMessage(USER_TYPE, message, data));
  }

}
