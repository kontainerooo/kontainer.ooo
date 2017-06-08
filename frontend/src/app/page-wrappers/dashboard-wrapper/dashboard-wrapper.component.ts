import { Component, OnInit } from '@angular/core';
import { GlobalDataService } from '../../services/global-data.service';

@Component({
  selector: 'kro-dashboard-wrapper',
  templateUrl: './dashboard-wrapper.component.html',
  styleUrls: ['./dashboard-wrapper.component.scss']
})
export class DashboardWrapperComponent implements OnInit {
  private loaded: boolean;

  constructor(private gds: GlobalDataService) {
    this.loaded = false;
  }

  ngOnInit() {
    this.gds.setAndGetUserAutomatically().subscribe(
      value => {
        this.loaded = true;
      },
      error => {
        console.log(error);
      }
    ); 
  }

}
