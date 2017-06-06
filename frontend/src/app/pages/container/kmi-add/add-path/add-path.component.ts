import { Component, OnInit } from '@angular/core';
import {
  FormGroup,
  Validators,
  FormBuilder
} from '@angular/forms';
import { MdlDialogReference, MdlSnackbarService } from 'angular2-mdl';
import { GlobalDataService } from '../../../../services/global-data.service';

@Component({
  selector: 'kro-add-path',
  templateUrl: './add-path.component.html',
  styleUrls: ['./add-path.component.scss']
})
export class AddPathComponent implements OnInit {
  private addPath: FormGroup;

  constructor(private mdldr: MdlDialogReference, private mdlss: MdlSnackbarService, private fb: FormBuilder, private gds: GlobalDataService) {

  }

  ngOnInit() {
    this.addPath = this.fb.group({
      path: ['', Validators.required]
    });
  }

  submit() {
    this.mdlss.showToast('Adding KMI...');
    this.gds.addKMI(this.addPath.get('path').value).subscribe(
      value => {
        this.mdldr.hide();
        this.mdlss.showToast(`Added KMI ${value} successfully!`);
      },
      error => {
        console.log(error);
      }
    )
  }

}
