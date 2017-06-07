/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { ContainerService } from './container.service';

describe('ContainerService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [ContainerService]
    });
  });

  it('should ...', inject([ContainerService], (service: ContainerService) => {
    expect(service).toBeTruthy();
  }));
});
