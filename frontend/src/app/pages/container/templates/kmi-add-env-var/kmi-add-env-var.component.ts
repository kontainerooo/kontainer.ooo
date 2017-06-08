import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

@Component({
  selector: 'kro-kmi-add-env-var',
  templateUrl: './kmi-add-env-var.component.html',
  styleUrls: ['./kmi-add-env-var.component.scss']
})
export class KmiAddEnvVarComponent implements OnInit {
  private env: FormGroup;

  constructor(private fb: FormBuilder) { }

  ngOnInit() {
    this.env = this.fb.group({
      key: ['', Validators.required],
      value: ['', Validators.required]
    });
  }

}
