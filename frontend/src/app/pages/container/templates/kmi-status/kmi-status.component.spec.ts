/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By, BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { DebugElement } from '@angular/core';
import { MdlModule } from 'angular2-mdl';

import { KmiStatusComponent } from './kmi-status.component';
import { KmiService } from '../../../../services/kmi.service';
import { SocketService } from '../../../../services/socket.service';

describe('KmiStatusComponent', () => {
  let component: KmiStatusComponent;
  let fixture: ComponentFixture<KmiStatusComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ KmiStatusComponent ],
      imports: [
        MdlModule,
        BrowserModule,
        FormsModule
      ],
      providers: [
        KmiService,
        SocketService
      ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(KmiStatusComponent);
    component = fixture.componentInstance;
    // fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
