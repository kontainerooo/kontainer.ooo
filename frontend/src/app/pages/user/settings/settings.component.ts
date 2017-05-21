import { Component, OnInit } from '@angular/core';
import {
  FormGroup,
  FormControl,
  Validators,
  FormBuilder
} from '@angular/forms';
import { DISABLE_NATIVE_VALIDITY_CHECKING } from 'angular2-mdl';

const emailValidator = Validators.pattern(/^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/);
const numberValidator = Validators.pattern(/\d+/);

@Component({
  selector: 'kio-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.scss'],
  providers: [
    {
      provide: DISABLE_NATIVE_VALIDITY_CHECKING,
      useValue: true
    }
  ]
})
export class SettingsComponent implements OnInit {
  form: FormGroup;

  constructor(private fb: FormBuilder) {

  }

  ngOnInit() {
    this.form = this.fb.group({
      firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      email: ['', Validators.compose([Validators.required, emailValidator])],
      username: ['', Validators.required],
      password: [''],
      rpassword: [''],
      phone: [''],
      country: ['', Validators.required],
      street: ['', Validators.required],
      housenumber: ['', Validators.compose([Validators.required, numberValidator])],
      additional: [''],
      zip: ['', Validators.compose([Validators.required, numberValidator])],
      city: ['', Validators.required]
    });

    this.form.valueChanges
      .subscribe((formValues) => {
        // TODO connect
      });
  }

  isFieldValid(name: string) {
    return !this.form.get(name).valid && this.form.get(name).touched;
  }

}
