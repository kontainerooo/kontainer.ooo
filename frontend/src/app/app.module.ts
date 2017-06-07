import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { MaterialModule } from '@angular/material';
import { RouterModule, Routes } from '@angular/router';
import { DashboardRoutingModule } from './page-wrappers/dashboard-wrapper/dashboard-routing.module';
import { appRouting } from './app-routing/app.routing';
import 'hammerjs';

import { GlobalDataService } from './services/global-data.service';

import { SocketService } from './services/socket.service';
import { UserService } from './services/user.service';
import { KenTheGuruService } from './services/ken-the-guru.service';
import { KmiService } from './services/kmi.service';
import { ContainerService } from './services/container.service';
import { ModuleService } from './services/module.service';

import { AppComponent } from './app.component';
import { SignInComponent } from './sign-in/sign-in.component';
import { PageNotFoundComponent } from './errors/page-not-found/page-not-found.component';
import { GatewayComponent } from './page-wrappers/gateway/gateway.component';
import { DashboardWrapperComponent } from './page-wrappers/dashboard-wrapper/dashboard-wrapper.component';
import { SidebarNavigationComponent } from './sidebar-navigation/sidebar-navigation.component';
import { DashboardToolbarComponent } from './dashboard-toolbar/dashboard-toolbar.component';
import { AccountComponent } from './sidebar-navigation/account/account.component';
import { AvatarComponent } from './sidebar-navigation/account/avatar/avatar.component';

@NgModule({
  declarations: [
    AppComponent,
    SignInComponent,
    PageNotFoundComponent,
    GatewayComponent,
    DashboardWrapperComponent,
    SidebarNavigationComponent,
    DashboardToolbarComponent,
    AccountComponent,
    AvatarComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    ReactiveFormsModule,
    HttpModule,
    MaterialModule.forRoot(),
    appRouting,
    DashboardRoutingModule
  ],
  providers: [
    GlobalDataService,
    SocketService,
    UserService,
    KenTheGuruService,
    KmiService,
    ContainerService,
    ModuleService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
