import { Component, OnInit, Inject } from '@angular/core';
import {
  FormGroup,
  Validators,
  FormBuilder
} from '@angular/forms';
import { MdlDialogReference, MdlSnackbarService } from 'angular2-mdl';
import { GlobalDataService } from '../../../../services/global-data.service';

@Component({
  selector: 'kro-add-container-modal',
  templateUrl: './add-container-modal.component.html',
  styleUrls: ['./add-container-modal.component.scss']
})
export class AddContainerModalComponent implements OnInit {
  private addContainer: FormGroup;

  constructor(@Inject('kmiId') private kmiId: number, private mdldr: MdlDialogReference, private mdlss: MdlSnackbarService, private fb: FormBuilder, private gds: GlobalDataService) {
    
  }

  ngOnInit() {
    this.addContainer = this.fb.group({
      name: ['', Validators.required]
    });
  }

  submit() {
    this.mdlss.showToast('Creating Container...');
    this.gds.createContainer(this.kmiId, this.addContainer.get('name').value).subscribe(
      value => {
        this.mdldr.hide();
        this.mdlss.showToast(`Created Container ${value} successfully!`);
      },
      error => {
        console.log(error);
      }
    )
  }

}
