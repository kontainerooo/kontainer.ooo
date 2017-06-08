import { Component, OnInit } from '@angular/core';
import { MdlSnackbarService } from 'angular2-mdl';
import { GlobalDataService } from '../../../../services/global-data.service';

@Component({
  selector: 'kro-kmi-install',
  templateUrl: './kmi-install.component.html',
  styleUrls: ['./kmi-install.component.scss']
})
export class KmiInstallComponent implements OnInit {

  constructor(private gds: GlobalDataService, private mdlss: MdlSnackbarService) {

  }

  ngOnInit() {

  }

  submit() {
    this.gds.sendCommand('install').subscribe(
      value => {
        console.log(value);
        this.mdlss.showToast('Installed!');
      },
      error => {
        console.log(error);
      }
    );
  }

}
