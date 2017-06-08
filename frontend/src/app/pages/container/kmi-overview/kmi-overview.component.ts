import { Component, OnInit } from '@angular/core';
import { container, kmi } from '../../../../messages/messages';
import { GlobalDataService } from '../../../services/global-data.service';

@Component({
  selector: 'kro-kmi-overview',
  templateUrl: './kmi-overview.component.html',
  styleUrls: ['./kmi-overview.component.scss']
})
export class KmiOverviewComponent implements OnInit {
  containers: Array<container.container>;

  constructor(private gds: GlobalDataService) {
    this.containers = [];
  }

  ngOnInit() {
    this.gds.getContainers().subscribe(
      value => {
        this.containers = value;
      },
      error => {
        console.log(error);
      }
    )
  }

  convertType(type: kmi.Type): string {
    switch(type) {
      case 0:
        return 'Webserver';
      default:
        return 'Default';
    }
  }

}
