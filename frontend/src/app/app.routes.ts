import { Routes, RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { HomeComponent } from './home/home.component';
import { MindMapComponent } from './mind-map/mind-map.component';

export const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'mind-map', component: MindMapComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
