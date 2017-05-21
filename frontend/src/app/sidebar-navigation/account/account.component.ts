import { Component, Input, OnChanges, ElementRef, Renderer } from '@angular/core';

import { SideNavElement } from '../../interfaces/side-nav-element';

@Component({
  selector: 'kro-account',
  templateUrl: './account.component.html',
  styleUrls: ['./account.component.scss']
})
export class AccountComponent implements OnChanges {
  @Input('bg') backgroundImage: string;
  @Input() realname: string;
  @Input() username: string;

  navigation: SideNavElement[];

  constructor(private elRef: ElementRef, private renderer: Renderer) {
    this.navigation = [
      {
        title: "Profile Settings",
        icon: "settings",
        route: "/user/settings"
      },
      {
        title: "Logout",
        icon: "exit_to_app",
        route: "/sign-in"
      }
    ];
  }

  ngOnChanges() {
    this.setBackground();
  }

  setBackground() {
    this.renderer.setElementStyle(this.elRef.nativeElement, 'background-image', `url('${this.backgroundImage}')`);
  }

}
