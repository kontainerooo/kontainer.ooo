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
import { KmiAddComponent } from '../../pages/container/kmi-add/kmi-add.component';

import { SearchKmiPipe } from '../../pipes/search-kmi.pipe';

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
      },
      {
        path: 'container/add',
        component: KmiAddComponent
      }
    ]
  }
];

@NgModule({
  declarations: [
    DashboardComponent,
    SettingsComponent,
    KmiOverviewComponent,
    KmiAddComponent,
    SearchKmiPipe
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
