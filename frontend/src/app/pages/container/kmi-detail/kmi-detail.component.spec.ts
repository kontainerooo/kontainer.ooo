/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';
import { MdlModule } from 'angular2-mdl';
import { ActivatedRoute, ActivatedRouteStub } from '../../../../testing/router-stubs';

import { KmiDetailComponent } from './kmi-detail.component';
import { KmiStatusComponent } from '../templates/kmi-status/kmi-status.component';
import { KmiService } from '../../../services/kmi.service';
import { SocketService } from '../../../services/socket.service';

describe('KmiDetailComponent', () => {
  let component: KmiDetailComponent;
  let fixture: ComponentFixture<KmiDetailComponent>;
  let activeRoute: ActivatedRouteStub;

  beforeEach(async(() => {
    activeRoute = new ActivatedRouteStub();

    TestBed.configureTestingModule({
      declarations: [
        KmiDetailComponent,
        KmiStatusComponent
      ],
      imports: [
        MdlModule
      ],
      providers: [
        KmiService,
        SocketService,
        {
          provide: ActivatedRoute,
          useValues: activeRoute
        }
      ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(KmiDetailComponent);
    component = fixture.componentInstance;

    activeRoute.testParams = {id: 1};

    fixture.detectChanges();
  });

  // TODO Tests

  // it('should create', () => {
  //   expect(component).toBeTruthy();
  // });
  
  // it('should get the correct container id', () => {
  //   expect(component.routeId).toEqual(1);
  // });
});
