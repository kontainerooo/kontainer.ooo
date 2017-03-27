/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { SocketService } from './socket.service';

describe('SocketService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [SocketService]
    });
  });

  it('should inject', inject([SocketService], (service: SocketService) => {
    expect(service).toBeTruthy();
  }));
});
