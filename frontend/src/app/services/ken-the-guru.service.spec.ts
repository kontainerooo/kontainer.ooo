/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { SocketService } from './socket.service';
import { KenTheGuruService } from './ken-the-guru.service';

describe('UserService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        SocketService,
        KenTheGuruService
      ]
    });
  });

  it('should inject', inject([KenTheGuruService], (service: KenTheGuruService) => {
    expect(service).toBeTruthy();
  }));
});
