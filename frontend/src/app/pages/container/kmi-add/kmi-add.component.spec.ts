/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';

import { KmiAddComponent } from './kmi-add.component';

describe('KmiAddComponent', () => {
  let component: KmiAddComponent;
  let fixture: ComponentFixture<KmiAddComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ KmiAddComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(KmiAddComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
