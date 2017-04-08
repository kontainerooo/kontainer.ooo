import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { MaterialModule } from '@angular/material';
import { DashboardRoutingModule } from './page-wrappers/dashboard-wrapper/dashboard-routing.module';
import { RouterModule, Routes } from '@angular/router';
import 'hammerjs';

import { AppComponent } from './app.component';
import { SignInComponent } from './sign-in/sign-in.component';
import { PageNotFoundComponent } from './errors/page-not-found/page-not-found.component';
import { GatewayComponent } from './page-wrappers/gateway/gateway.component';
import { DashboardWrapperComponent } from './page-wrappers/dashboard-wrapper/dashboard-wrapper.component';
import { SidebarNavigationComponent } from './sidebar-navigation/sidebar-navigation.component';
import { DashboardToolbarComponent } from './dashboard-toolbar/dashboard-toolbar.component';
import { AccountComponent } from './sidebar-navigation/account/account.component';
import { AvatarComponent } from './sidebar-navigation/account/avatar/avatar.component';

const appRoutes: Routes = [
  {
    path: 'sign-in',
    component: SignInComponent
  },
  // { path: 'hero/:id',      component: HeroDetailComponent },
  // {
  //   path: 'heroes',
  //   component: HeroListComponent,
  //   data: { title: 'Heroes List' }
  // },
  {
    path: '',
    redirectTo: '/dashboard',
    pathMatch: 'full'
  },
  {
    path: '**',
    component: PageNotFoundComponent
  }
];

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
    HttpModule,
    MaterialModule.forRoot(),
    RouterModule.forRoot(appRoutes),
    DashboardRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
