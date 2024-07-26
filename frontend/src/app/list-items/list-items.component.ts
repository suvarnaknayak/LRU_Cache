import { Component, Input } from '@angular/core';
import { CacheService } from '../services/cache.service';
import { Observable, Subject } from 'rxjs';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-list-items',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './list-items.component.html',
  styleUrl:'./list-items.components.scss'
})
export class ListItemsComponent {

  @Input('update') update!: Observable<void>

  list: Record<string, string> = {}

  constructor(private cacheService: CacheService) { }

  ngOnInit(): void {
    this.update?.subscribe(this.getKeyValuePairs.bind(this))

    this.getKeyValuePairs()
  }

  getKeyValuePairs(): void {

    this.cacheService.getCacheList().subscribe({
      next: (data) => {
        this.list = data as typeof this.list
        console.log("ggvbhjn",this.list)
      },
      error: (error) => {
        console.log(error);
      }
    });
  }

  delete(key:string){
    this.cacheService.deleteValue(key).subscribe({
      next: (data) => {
        console.log(data)
        this.getKeyValuePairs()
      },
      error: (error) => {
        console.log(error);
      }
    });
  }
}
