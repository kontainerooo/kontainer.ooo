/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { SocketService } from './socket.service';
import { UserService } from './user.service';

describe('UserService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        SocketService,
        UserService
      ]
    });
  });

  it('should inject', inject([UserService], (service: UserService) => {
    expect(service).toBeTruthy();
  }));
});
