import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { RouterModule, Routes } from '@angular/router';
import { MaterialModule } from '@angular/material';
import { MdlModule } from 'angular2-mdl';

import { DashboardWrapperComponent } from './dashboard-wrapper.component';
import { DashboardComponent } from '../../pages/dashboard/dashboard.component';
import { SettingsComponent } from '../../pages/user/settings/settings.component';
import { KmiOverviewComponent } from '../../pages/container/kmi-overview/kmi-overview.component';

const dashboardRoutes = [
  {
    path: '',
    component: DashboardWrapperComponent,
    children: [
      {
        path: '',
        component: DashboardComponent
      },
      {
        path: 'dashboard',
        component: DashboardComponent
      },
      {
        path: 'user/settings',
        component: SettingsComponent
      },
      {
        path: 'container',
        component: KmiOverviewComponent
      }
    ]
  }
];

@NgModule({
  declarations: [
    DashboardComponent,
    SettingsComponent,
    KmiOverviewComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    RouterModule.forChild(dashboardRoutes),
    MaterialModule,
    MdlModule
  ],
  exports: [
    RouterModule
  ]
})
export class DashboardRoutingModule { }
