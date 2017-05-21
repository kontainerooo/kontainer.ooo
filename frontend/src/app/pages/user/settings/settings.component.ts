import { Component, OnInit } from '@angular/core';
import {
  FormGroup,
  FormControl,
  Validators,
  FormBuilder
} from '@angular/forms';
import { DISABLE_NATIVE_VALIDITY_CHECKING } from 'angular2-mdl';
import { UserService } from '../../../services/user.service';

const emailValidator = Validators.pattern(/^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/);
const numberValidator = Validators.pattern(/\d+/);

@Component({
  selector: 'kio-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.scss'],
  providers: [
    UserService,
    {
      provide: DISABLE_NATIVE_VALIDITY_CHECKING,
      useValue: true
    }
  ]
})
export class SettingsComponent implements OnInit {
  form: FormGroup;

  constructor(private fb: FormBuilder, private us: UserService) {
    
  }

  ngOnInit() {
    this.us.messages.subscribe(
      (value) => {
        console.log(value);
      },
      (error) => {
        console.log(error);
      }
    );

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

  onSubmit() {
    this.us.next('CreateUserRequest', {
      username: this.form.get('username'),
      config: {
        admin: true, // TODO change to something more secure
        email: this.form.get('email'),
        password: this.form.get('password'),
        salt: `512 bytes°random string`,
        address: {
          postcode: this.form.get('zip'),
          city: this.form.get('city'),
          country: this.form.get('country'),
          street: this.form.get('street'),
          houseno: this.form.get('housenumber'),
          additional: this.form.get('additional')
        },
        phone: this.form.get('phone'),
        image: 'http://static.giantbomb.com/uploads/scale_small/0/9517/2816097-tumblr_n45cr8dmj61ty0km0o7_1280.png'
      }
    })
  }

}
