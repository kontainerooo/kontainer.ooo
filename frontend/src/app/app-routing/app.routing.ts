import { ModuleWithProviders } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { SignInComponent } from '../sign-in/sign-in.component';
import { SettingsComponent } from '../pages/user/settings/settings.component';
import { PageNotFoundComponent } from '../errors/page-not-found/page-not-found.component';

const appRoutes: Routes = [
  {
    path: 'sign-in',
    component: SignInComponent
  },
  {
    path: 'register',
    component: SettingsComponent
  },
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
export const appRouting: ModuleWithProviders = RouterModule.forRoot(appRoutes);
