import { Injectable } from '@angular/core';
import { Observable, Subject, Subscription } from 'rxjs/Rx';
import { PartialObserver } from 'rxjs/Observer';
import { SocketService } from './socket.service'

const SOCKET_URL = 'ws://localhost:8083';
const USER_TYPE = 'kmi';

@Injectable()
export class KmiService {
  public messages: Subject<Uint8Array>;

  constructor(private wsService: SocketService) {
    this.messages = <Subject<Uint8Array>>this.wsService
      .connect(SOCKET_URL)
      .map((response: MessageEvent): object => {
        return this.wsService.decodeMessage(new Uint8Array(response.data), USER_TYPE);
      });
  }

  public next(message: string, data: object) {
    this.messages.next(this.wsService.encodeMessage(USER_TYPE, message, data));
  }

}
