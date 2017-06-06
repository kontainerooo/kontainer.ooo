import { Component, OnInit } from '@angular/core';
import { MdlDialogService } from 'angular2-mdl';
import { kmi } from '../../../../messages/messages';
import { AddPathComponent } from './add-path/add-path.component';
import { GlobalDataService } from '../../../services/global-data.service';

@Component({
  selector: 'kro-kmi-add',
  templateUrl: './kmi-add.component.html',
  styleUrls: ['./kmi-add.component.scss']
})
export class KmiAddComponent implements OnInit {
  kmiModules: Array<kmi.KMDI>;

  constructor(private mdlds: MdlDialogService, private gds: GlobalDataService) {
    this.kmiModules = [];
  }

  ngOnInit() {
    this.gds.setAndGetAvailableKMI().subscribe(
      value => {
        this.kmiModules = value;
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

  showAddKMI() {
    let addDialog = this.mdlds.showCustomDialog({
      component: AddPathComponent,
      providers: [],
      isModal: true,
      styles: {'width': '350px'},
      clickOutsideToClose: true,
      enterTransitionDuration: 400,
      leaveTransitionDuration: 400
    });
  }

}
