/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By, BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { DebugElement } from '@angular/core';
import { MdlModule } from 'angular2-mdl';

import { KmiAddComponent } from './kmi-add.component';
import { SearchKmiPipe } from '../../../pipes/search-kmi.pipe';

describe('KmiAddComponent', () => {
  let component: KmiAddComponent;
  let fixture: ComponentFixture<KmiAddComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        KmiAddComponent,
        SearchKmiPipe
      ],
      imports: [
        BrowserModule,
        FormsModule,
        MdlModule
      ]
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
