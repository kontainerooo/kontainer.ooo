/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';

import { MaterialModule } from '@angular/material';

import { AccountComponent } from './account.component';

import { RouterLinkStubDirective } from '../../../testing/router-stubs';

describe('AccountComponent', () => {
  let component: AccountComponent;
  let fixture: ComponentFixture<AccountComponent>;

  let linkDes: DebugElement[];
  let links: RouterLinkStubDirective[];

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
          AccountComponent,
          RouterLinkStubDirective
      ],
      imports: [
          MaterialModule
      ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AccountComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();

    linkDes = fixture.debugElement
      .queryAll(By.directive(RouterLinkStubDirective));
    links = linkDes
      .map(de => de.injector.get(RouterLinkStubDirective) as RouterLinkStubDirective);
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should route correctly', () => {
    it('should get all router links from template', () => {
      expect(links.length).toBe(1, 'should have 1 link');
      expect(links[0].linkParams).toBe('/sign-in', '1st link should go to SignIn');
    });

    it('should route to the sign-in on logout item', () => {
      checkRouteClick(linkDes[0], links[0], '/sign-in');
    });
  });

  function checkRouteClick(linkDe: DebugElement, link: RouterLinkStubDirective, route: string) {
    expect(link.navigatedTo).toBeNull('link should not have navigated yet');

    linkDe.triggerEventHandler('click', null);
    fixture.detectChanges();

    expect(link.navigatedTo).toBe(route);
  }
});
