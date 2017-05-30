import { Component, OnInit, OnDestroy } from '@angular/core';
import {
  FormGroup,
  Validators,
  FormBuilder
} from '@angular/forms';
import { ActivatedRoute } from '@angular/router';
import { DISABLE_NATIVE_VALIDITY_CHECKING, MdlSnackbarService } from 'angular2-mdl';
import { UserService } from '../../../services/user.service';
import { GlobalDataService } from '../../../services/global-data.service';
import { ProtoResponse } from '../../../interfaces/proto-response';
import { user } from '../../../../messages/messages';

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
export class SettingsComponent implements OnInit, OnDestroy {
  form: FormGroup;
  edit: boolean;
  loading: boolean;

  constructor(private fb: FormBuilder, private us: UserService, private ar: ActivatedRoute, private mdlss: MdlSnackbarService, private gds: GlobalDataService) {
    this.loading = true;
    if(ar.snapshot.url[0].path == 'register') {
      this.edit = false;
    } else {
      this.edit = true;
    }
  }

  ngOnInit() {
    this.form = this.fb.group({
      firstName: [{value: '', disabled: this.edit}, Validators.required],
      lastName: [{value: '', disabled: this.edit}, Validators.required],
      email: ['', Validators.compose([Validators.required, emailValidator])],
      username: [{value: '', disabled: this.edit}, Validators.required],
      password: [''],
      rpassword: [''],
      phone: [''],
      country: ['', Validators.required],
      street: ['', Validators.required],
      housenumber: ['', Validators.compose([Validators.required, numberValidator])],
      additional: [''],
      zip: ['', Validators.required],
      city: ['', Validators.required]
    });

    if(this.edit) {
      this.gds.setAndGetUserById(3).subscribe(
        user => {
          this.gotUser(user);
        },
        error => {
          console.log(error);
        }
      );
    }
  }

  ngOnDestroy() {
    this.us.messages.unsubscribe();
  }

  isFieldValid(name: string) {
    return !this.form.get(name).valid && this.form.get(name).touched;
  }

  onSubmit() {
    this.mdlss.showToast('Saving...');

    let testObject: user.User$Properties = {
      username: this.form.get('username').value,
      config: {
        admin: true, // TODO change to something more secure
        email: this.form.get('email').value,
        password: this.form.get('password').value,
        address: {
          postcode: this.form.get('zip').value,
          city: this.form.get('city').value,
          country: this.form.get('country').value,
          street: this.form.get('street').value,
          houseno: this.form.get('housenumber').value,
          additional: this.form.get('additional').value
        },
        phone: this.form.get('phone').value,
        image: 'http://static.giantbomb.com/uploads/scale_small/0/9517/2816097-tumblr_n45cr8dmj61ty0km0o7_1280.png'
      }
    };
    
    if(!this.edit) {
      this.gds.registerUser(testObject).subscribe(
        id => {
          this.mdlss.showToast('User saved');
        },
        error => {
          console.log(error);
        }
      );
    } else {
      this.gds.editUser(3, testObject.config).subscribe(
        success => {
          if(success) {
            this.mdlss.showToast('User saved');
          }
        },
        error => {
          console.log(error);
        }
      );
    }
  }

  gotUser(user: user.User) {
    if(user.config) {
      if(user.config.email) {
        this.form.get('email').setValue(user.config.email);
      }
      if(user.username) {
        this.form.get('username').setValue(user.username);
      }
      if(user.config.phone) {
        this.form.get('phone').setValue(user.config.phone);
      }
      if(user.config.address) {
        if(user.config.address.country) {
          this.form.get('country').setValue(user.config.address.country);
        }
        if(user.config.address.street) {
          this.form.get('street').setValue(user.config.address.street);
        }
        if(user.config.address.houseno) {
          this.form.get('housenumber').setValue(user.config.address.houseno);
        }
        if(user.config.address.additional) {
          this.form.get('additional').setValue(user.config.address.additional);
        }
        if(user.config.address.postcode) {
          this.form.get('zip').setValue(user.config.address.postcode);
        }
        if(user.config.address.city) {
          this.form.get('city').setValue(user.config.address.city);
        }
      }
    }
    this.loading = false;
  }

}
