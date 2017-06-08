/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';

import { MaterialModule } from '@angular/material';

import { DashboardWrapperComponent } from './dashboard-wrapper.component';
import { DashboardToolbarComponent } from '../../dashboard-toolbar/dashboard-toolbar.component';
import { SidebarNavigationComponent } from '../../sidebar-navigation/sidebar-navigation.component';
import { AccountComponent } from '../../sidebar-navigation/account/account.component';
import { AvatarComponent } from '../../sidebar-navigation/account/avatar/avatar.component';

import { RouterLinkStubDirective, RouterOutletStubComponent } from '../../../testing/router-stubs';

describe('DashboardWrapperComponent', () => {
  let component: DashboardWrapperComponent;
  let fixture: ComponentFixture<DashboardWrapperComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
          DashboardWrapperComponent,
          DashboardToolbarComponent,
          SidebarNavigationComponent,
          AccountComponent,
          AvatarComponent,
          RouterLinkStubDirective,
          RouterOutletStubComponent
      ],
      imports: [
          MaterialModule
      ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DashboardWrapperComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  // it('should create', () => {
  //   expect(component).toBeTruthy();
  // });
});
