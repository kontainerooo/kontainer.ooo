import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { MdlSnackbarService } from 'angular2-mdl';
import { GlobalDataService } from '../../../../services/global-data.service';

const template = {
	"name": "upload",
	"parameters": [
		{
			"name": "title",
			"type": "value"
		},
		{
			"name": "dir",
			"type": "value"
		},
		{
			"name": "list",
			"type": "value"
		}
	]
};

@Component({
  selector: 'kro-kmi-upload-file',
  templateUrl: './kmi-upload-file.component.html',
  styleUrls: ['./kmi-upload-file.component.scss']
})
export class KmiUploadFileComponent implements OnInit {
  private title: string;
  private upload: FormGroup;
  private file;

  constructor(private gds: GlobalDataService, private fb: FormBuilder, private mdlss: MdlSnackbarService) {
    this.title = this.gds.getValueSnapshot('upload', 'title');
  }

  ngOnInit() {
    this.upload = this.fb.group({
      path: ['', Validators.required],
      file: ['']
    });
  }

  onFileChange(event: Event) {
    this.file = (<HTMLInputElement>event.srcElement).files[0];
  }

  submit() {
    let fr = new FileReader();
    let resultArray: Uint8Array;
    fr.onload = () => {
      let ab = fr.result;
      resultArray = new Uint8Array(ab);
      this.gds.uploadFile(this.upload.get('path').value, resultArray, true).subscribe(
        value => {
          if(value) {
            this.mdlss.showToast('File uploaded!');
          }
        },
        error => {
          console.log(error);
        }
      );
    };
    fr.readAsArrayBuffer(this.file);
  }

}
