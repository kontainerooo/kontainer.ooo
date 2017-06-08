/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';

import { KmiInstallComponent } from './kmi-install.component';

describe('KmiInstallComponent', () => {
  let component: KmiInstallComponent;
  let fixture: ComponentFixture<KmiInstallComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ KmiInstallComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(KmiInstallComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  // it('should create', () => {
  //   expect(component).toBeTruthy();
  // });
});
