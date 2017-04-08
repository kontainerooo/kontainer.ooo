import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { MaterialModule } from '@angular/material';

import { DashboardWrapperComponent } from './dashboard-wrapper.component';
import { DashboardComponent } from '../../pages/dashboard/dashboard.component';
import { SettingsComponent } from '../../pages/user/settings/settings.component';

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
      }
    ]
  }
];

@NgModule({
  declarations: [
    DashboardComponent,
    SettingsComponent
  ],
  imports: [
    RouterModule.forChild(dashboardRoutes),
    MaterialModule
  ],
  exports: [
    RouterModule
  ]
})
export class DashboardRoutingModule { }
