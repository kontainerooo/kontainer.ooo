import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { MaterialModule } from '@angular/material';
import { DashboardModule } from './page-wrappers/dashboard/dashboard.module';
import { RouterModule, Routes } from '@angular/router';
import 'hammerjs';

import { AppComponent } from './app.component';
import { SignInComponent } from './sign-in/sign-in.component';
import { PageNotFoundComponent } from './errors/page-not-found/page-not-found.component';
import { GatewayComponent } from './page-wrappers/gateway/gateway.component';
import { DashboardComponent } from './page-wrappers/dashboard/dashboard.component';
import { SidebarNavigationComponent } from './sidebar-navigation/sidebar-navigation.component';
import { DashboardToolbarComponent } from './dashboard-toolbar/dashboard-toolbar.component';
import { AccountComponent } from './sidebar-navigation/account/account.component';
import { AvatarComponent } from './sidebar-navigation/account/avatar/avatar.component';
import { SettingsComponent } from './pages/user/settings/settings.component';
import { KioCardFullWidthDirective } from './directives/kio-card.directive';

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
    DashboardComponent,
    SidebarNavigationComponent,
    DashboardToolbarComponent,
    AccountComponent,
    AvatarComponent,
    SettingsComponent,
    KioCardFullWidthDirective
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    MaterialModule.forRoot(),
    RouterModule.forRoot(appRoutes),
    DashboardModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
