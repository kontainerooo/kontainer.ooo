import { Component, OnInit } from '@angular/core';

import { SideNavElement } from '../interfaces/side-nav-element';

@Component({
  selector: 'kio-sidebar-navigation',
  templateUrl: './sidebar-navigation.component.html',
  styleUrls: ['./sidebar-navigation.component.scss']
})
export class SidebarNavigationComponent implements OnInit {
  navigation: SideNavElement[];

  constructor() {
    this.navigation = [
      {
        title: "Dashboard",
        icon: "dashboard",
        route: "/dashboard"
      },
      {
        title: "Containers",
        icon: "view_module",
        route: "/dashboard"
      }
    ]
  }

  ngOnInit() {
  }

}
