/* tslint:disable:no-unused-variable */
import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';
import { DebugElement } from '@angular/core';

import { MaterialModule } from '@angular/material';
import 'hammerjs';

import { SidebarNavigationComponent } from './sidebar-navigation.component';
import { AccountComponent } from './account/account.component';
import { AvatarComponent } from './account/avatar/avatar.component';

import { RouterLinkStubDirective } from '../../testing/router-stubs';

describe('SidebarNavigationComponent', () => {
  let component: SidebarNavigationComponent;
  let fixture: ComponentFixture<SidebarNavigationComponent>;

  let linkDes: DebugElement[];
  let links: RouterLinkStubDirective[];

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
          SidebarNavigationComponent,
          AccountComponent,
          AvatarComponent,
          RouterLinkStubDirective
      ],
      imports: [
          MaterialModule
      ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SidebarNavigationComponent);
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
      expect(links.length).toBe(component.navigation.length, 'should have as links as the component navigation');

      for(let i in links) {
        expect(links[i].linkParams).toBe(component.navigation[i].route, `link ${i} should go to SignIn`);
      }
    });

    it('should route to the component defined sites', () => {
      for(let i in links) {
        checkRouteClick(linkDes[i], links[i], component.navigation[i].route);
      }
    });
  });

  function checkRouteClick(linkDe: DebugElement, link: RouterLinkStubDirective, route: string) {
    expect(link.navigatedTo).toBeNull('link should not have navigated yet');

    linkDe.triggerEventHandler('click', null);
    fixture.detectChanges();

    expect(link.navigatedTo).toBe(route);
  }
});
