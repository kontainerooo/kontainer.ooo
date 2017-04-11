/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By, BrowserModule} from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { DebugElement } from '@angular/core';
import { MdlModule } from 'angular2-mdl';

import { KmiOverviewComponent } from './kmi-overview.component';
import { RouterLinkStubDirective } from '../../../../testing/router-stubs'

describe('KmiOverviewComponent', () => {
  let component: KmiOverviewComponent;
  let fixture: ComponentFixture<KmiOverviewComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        KmiOverviewComponent,
        RouterLinkStubDirective
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
    fixture = TestBed.createComponent(KmiOverviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
