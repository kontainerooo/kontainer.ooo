import { Injectable } from '@angular/core';
import { GlobalData } from '../interfaces/global-data';
import { UserService } from './user.service';
import { user } from '../../messages/messages';
import { ProtoResponse } from '../interfaces/proto-response';
import { Observable, ReplaySubject } from 'rxjs/Rx';

@Injectable()
export class GlobalDataService {
  private gd: GlobalData;

  constructor(private us: UserService) {
    this.gd = {};
  }
  
  getUserId(): number {
    if(this.gd.user && this.gd.user.ID) {
      return this.gd.user.ID;
    } else {
      return -1;
    }
  }

  getUserSnapshot(): user.User {
    if(this.gd.user) {
      return this.gd.user;
    } else {
      return undefined;
    }
  }

  setAndGetUserById(id: number): Observable<user.User> {
    let obs = this.us.messages
      .share()
      .first((value: ProtoResponse) => {
        return value.message == 'GetUserResponse';
      })
      .map((value: ProtoResponse): user.User => {
        return user.User.create(user.GetUserResponse.from(value.data).user);
      });

    obs.subscribe(
      user => {
        this.gd.user = user;
      },
      error => {
        console.log(error);
      }
    );
    
    this.us.next('GetUserRequest', {
      ID: id
    });

    return obs;
  }

  registerUser(userProperties: user.User$Properties): Observable<number> {
    let obs = this.us.messages
      .share()
      .first((value: ProtoResponse) => {
        return value.message == 'CreateUserResponse';
      })
      .map((value: ProtoResponse): number => {
        return user.CreateUserResponse.from(value.data).ID;
      });

    this.us.next('CreateUserRequest', user);

    return obs;
  }

  editUser(id: number, configProperties: user.Config$Properties): Observable<boolean> {
    let obs = this.us.messages
      .share()
      .first((value: ProtoResponse) => {
        return value.message == 'EditUserResponse';
      })
      .map((value: ProtoResponse): boolean => {
        return user.EditUserResponse.from(value.data).error ? false : true;
      });

    obs.subscribe(
      success => {
        if(success) {
          this.setAndGetUserById(this.getUserId());
        }
      },
      error => {
        console.log(error);
      }
    )

    this.us.next('EditUserRequest', {
      ID: id,
      config: configProperties
    });

    return obs;
  }

  // logIn(username: string, password: string): Promise<user.User> {
  //   this.us.next('CheckLoginCredentialsRequest')
  // }
}
