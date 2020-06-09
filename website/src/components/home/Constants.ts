import { SelectItem } from "../../entities/common/SelectItem";
import { CheckboxGroupItem } from "../../entities/common/CheckboxGroupItem";

export const laneSelectItems = [
  {
    label: 'Top',
    value: 'top',
  },
  {
    label: 'Jungle',
    value: 'jungle',
  },
  {
    label: 'Middle',
    value: 'middle',
  },
  {
    label: 'Bottom',
    value: 'bottom',
  },
  {
    label: 'Support',
    value: 'support',
  },
] as SelectItem[]

export const attributeCheckboxGroupItems = [
  {
    label: 'Kills',
    name: 'kills'
  },
  {
    label: 'Deaths',
    name: 'deaths'
  },
  {
    label: 'Assists',
    name: 'assists'
  },
  {
    label: 'Pick Percentage',
    name: 'pick_percentage'
  },
  {
    label: 'Physical Damage',
    name: 'physical_damage'
  },
  {
    label: 'Magic Damage',
    name: 'magic_damage'
  },
  {
    label: 'True Damage',
    name: 'true_damage'
  },
  {
    label: 'Total Damage',
    name: 'total_damage'
  },
  {
    label: 'Damage Taken',
    name: 'damage_taken'
  },
  {
    label: 'Healing',
    name: 'healing'
  },
  {
    label: 'Max Kill Spree',
    name: 'max_kill_spree'
  },
  {
    label: 'Gold',
    name: 'gold'
  },
  {
    label: 'Minions Killed',
    name: 'minions_killed'
  },
  {
    label: 'Jungle CS',
    name: 'jungle_cs'
  },
  {
    label: 'Enemy Jungle CS',
    name: 'enemy_jungle_cs'
  },
  {
    label: 'Team Jungle CS',
    name: 'team_jungle_cs'
  },
] as CheckboxGroupItem[]

export const distanceMethodSelectItems = [
  {
    label: 'Euclidean',
    value: 'euclidean',
  },
  {
    label: 'Manhattan',
    value: 'manhattan',
  },
] as SelectItem[]