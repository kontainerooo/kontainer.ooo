/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';

import { SignInComponent } from './sign-in.component';
import { GatewayComponent } from '../page-wrappers/gateway/gateway.component';
import { MaterialModule } from '@angular/material';

import { RouterLinkStubDirective } from '../../testing/router-stubs';

describe('SignInComponent', () => {
  let component: SignInComponent;
  let fixture: ComponentFixture<SignInComponent>;
  let de: DebugElement;
  let el: HTMLElement;

  let linkDes: DebugElement[];
  let links: RouterLinkStubDirective[];

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
          SignInComponent,
          GatewayComponent,
          RouterLinkStubDirective
      ],
      imports: [
          MaterialModule.forRoot()
      ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SignInComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    linkDes = fixture.debugElement
      .queryAll(By.directive(RouterLinkStubDirective));
    links = linkDes
      .map(de => de.injector.get(RouterLinkStubDirective) as RouterLinkStubDirective);
  });

  // it('should create', () => {
  //   expect(component).toBeTruthy();
  // });

  // it('should get all router links from template', () => {
  //   expect(links.length).toBe(1, 'should have 1 link');
  //   expect(links[0].linkParams).toBe('/dashboard', '1st link should go to Dashboard');
  // });

  // it('should route to the dashboard after clicking sign in', () => {
  //   const linkDe = linkDes[0];
  //   const link = links[0];

  //   expect(link.navigatedTo).toBeNull('link should not have navigated yet');

  //   linkDe.triggerEventHandler('click', null);
  //   fixture.detectChanges();

  //   expect(link.navigatedTo).toBe('/dashboard');
  // });
});
