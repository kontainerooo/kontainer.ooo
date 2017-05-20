import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { RouterModule, Routes } from '@angular/router';
import { MaterialModule } from '@angular/material';
import { MdlModule } from 'angular2-mdl';

import { SearchKmiPipe } from '../../pipes/search-kmi.pipe';

import { SocketService } from '../../services/socket.service';
import { KmiService } from '../../services/kmi.service';

import { DashboardWrapperComponent } from './dashboard-wrapper.component';
import { DashboardComponent } from '../../pages/dashboard/dashboard.component';
import { SettingsComponent } from '../../pages/user/settings/settings.component';
import { KmiOverviewComponent } from '../../pages/container/kmi-overview/kmi-overview.component';
import { KmiAddComponent } from '../../pages/container/kmi-add/kmi-add.component';
import { KmiDetailComponent } from './../../pages/container/kmi-detail/kmi-detail.component';
import { KmiStatusComponent } from './../../pages/container/templates/kmi-status/kmi-status.component';

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
      },
      {
        path: 'container/:id',
        component: KmiDetailComponent
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
    SearchKmiPipe,
    KmiDetailComponent,
    KmiStatusComponent
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
  ],
  providers: [
    SocketService,
    KmiService
  ]
})
export class DashboardRoutingModule { }
