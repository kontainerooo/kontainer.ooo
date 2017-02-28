/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';

import { MaterialModule } from '@angular/material';

import { DashboardComponent } from './dashboard.component';
import { DashboardToolbarComponent } from '../../dashboard-toolbar/dashboard-toolbar.component';
import { SidebarNavigationComponent } from '../../sidebar-navigation/sidebar-navigation.component';
import { AccountComponent } from '../../sidebar-navigation/account/account.component';
import { AvatarComponent } from '../../sidebar-navigation/account/avatar/avatar.component';

import { RouterLinkStubDirective } from '../../../testing/router-stubs';

describe('DashboardComponent', () => {
  let component: DashboardComponent;
  let fixture: ComponentFixture<DashboardComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
          DashboardComponent,
          DashboardToolbarComponent,
          SidebarNavigationComponent,
          AccountComponent,
          AvatarComponent,
          RouterLinkStubDirective
      ],
      imports: [
          MaterialModule
      ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DashboardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
