import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { MdlSnackbarService } from 'angular2-mdl';
import { GlobalDataService } from '../../../../services/global-data.service';

@Component({
  selector: 'kro-kmi-add-env-var',
  templateUrl: './kmi-add-env-var.component.html',
  styleUrls: ['./kmi-add-env-var.component.scss']
})
export class KmiAddEnvVarComponent implements OnInit {
  private env: FormGroup;
  private ip: string;

  constructor(private fb: FormBuilder, private gds: GlobalDataService, private mdlss: MdlSnackbarService) {
    this.ip = 'Not yet requested';
  }

  ngOnInit() {
    this.env = this.fb.group({
      name: ['', Validators.required],
      interface: ['', Validators.required]
    });
  }

  submit() {
    this.gds.setLink(this.env.get('name').value, this.env.get('interface').value).subscribe(
      value => {
        if(value) {
          this.mdlss.showToast('Link created!');
        }
      },
      error => {
        console.log(error);
      }
    );
  }

  getIP() {
    this.gds.sendCommand('getip').subscribe(
      value => {
        this.ip = value;
      },
      error => {
        console.log(error);
      }
    );
  }

}
