/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { SocketService } from './socket.service';
import { KmiService } from './kmi.service';

describe('KmiService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        SocketService,
        KmiService
      ]
    });
  });

  it('should inject', inject([KmiService], (service: KmiService) => {
    expect(service).toBeTruthy();
  }));
});
