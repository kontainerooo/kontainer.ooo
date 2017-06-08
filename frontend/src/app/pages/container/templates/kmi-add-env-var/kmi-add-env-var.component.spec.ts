/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';

import { KmiAddEnvVarComponent } from './kmi-add-env-var.component';

describe('KmiAddEnvVarComponent', () => {
  let component: KmiAddEnvVarComponent;
  let fixture: ComponentFixture<KmiAddEnvVarComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ KmiAddEnvVarComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(KmiAddEnvVarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
