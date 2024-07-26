import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { CreateCacheComponent } from "./create-cache/create-cache.component";
import { Subject } from 'rxjs';
import { ListItemsComponent } from "./list-items/list-items.component";

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, CreateCacheComponent, ListItemsComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'angular-frontend';

  private updateList = new Subject<void>()

  get $update(){
    return this.updateList.asObservable()
  }

  handleValueAdded() {
    this.updateList.next()
  }
}
