<script lang="ts" setup>
import {reactive} from 'vue'
import {
  CurrentTime,
  ScheduleDuration,
  IsScheduled,
  CancelSchedule,
  RemainingTime
} from '../../wailsjs/go/main/App'

const data = reactive({
  description: "Select the time or duration and click schedule to enable it.",
  resultText: "",
  currentTime: "",
  remainingTime: "",
  minutes: 60,
  scheduledMinutes: 0,
  action: "sleep",
  scheduled: false,
  buttonActive: function(minutes : number) {
    return {
      "btn-active": data.scheduledMinutes == minutes && data.scheduled,
    }
  },
  textFieldActive: function() {
    return {
      "txt-active": data.scheduled && data.minutes == data.scheduledMinutes,
    }
  }
})

function currentTime() {
  CurrentTime().then(result => {
    data.currentTime = result
  })
}

setInterval(currentTime, 500)

function remainingTime() {
  RemainingTime().then(result => {
    data.remainingTime = result
  })
}

setInterval(remainingTime, 500)

function checkIsScheduled() {
  IsScheduled().then(result => {
    data.scheduled = result
  })
}

setInterval(checkIsScheduled, 500)

function scheduleDuration(minutes : number, event : MouseEvent) {
  if (minutes != 0) {
    data.minutes = minutes
  }

  ScheduleDuration(`${data.minutes}m`).then(result => {
    data.scheduledMinutes = data.minutes
    data.resultText = result
    data.scheduled = true
  })
}

function cancel() {
  CancelSchedule().then(result => {
    data.resultText = result
    data.scheduled = false
    data.scheduledMinutes = 0
  })
}

</script>

<template>
  <main>
    <div id="description" class="description">{{ data.description }}</div>
    <div id="time" class="time">{{ data.currentTime }}</div>
    <div id="duration-shortcuts" class="duration-shortcuts">
      <input type="radio"
              class="btn-group btn-first"
              v-bind:class="data.buttonActive(10)"
              value="10"
              name="duration10" />
      <label for="duration10" @click="scheduleDuration(10, $event)">10min</label>
      <input type="radio"
              class="btn-group"
              v-bind:class="data.buttonActive(20)"
              value="20"
              name="duration20" />
      <label for="duration20" @click="scheduleDuration(20, $event)">20min</label>
      <input type="radio"
              class="btn-group"
              v-bind:class="data.buttonActive(30)"
              value="30"
              name="duration30" />
      <label for="duration30" @click="scheduleDuration(30, $event)">30min</label>
      <input type="radio"
              class="btn-group btn-last"
              v-bind:class="data.buttonActive(40)"
              value="40"
              name="duration40" />
      <label for="duration40" @click="scheduleDuration(40, $event)">40min</label>
    </div>
    <div id="duration-form" class="duration-form">
      <input name="minutes"
             type="number"
             class="txt"
             v-model="data.minutes"
             v-bind:class="data.textFieldActive()"
             pattern="[0-9]+"
             step="1"
             min="1"
             max="1440">
      <span class="unit">min</span>
      <button class="btn" @click="scheduleDuration(0, $event)">Schedule</button>
    </div>
    <div id="result" class="result">
      {{ data.resultText }}
    </div>
    <div id="remaining-time" class="remaining-time" v-if="data.scheduled">
      <span>{{ data.remainingTime }}</span>
      <button class="btn" @click="cancel">Cancel</button>
    </div>
  </main>
</template>

<style scoped>
.description {
  padding: 1rem;
  font-size: 1.25rem;
}

.time {
  padding: 1rem;
  font-size: 2rem;
}

.duration-form {
  padding: 1rem;
}

.description, .time, .duration-shortcuts, .duration-form {
  display: block;
}

.remaining-time span {
  font-size: 1.5rem;
}

.unit {
  margin-left:-20px;
}

.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.btn, .txt {
  color: white;
  padding: 1rem;
  margin: 0 1.5rem;
  background-color: black;
  border: none;
  border-radius: 0.25rem;
  cursor: pointer;
  font-size: 1rem;
}

.btn.btn-active {
  background-color: #8155ba;
  color: white;
}

.txt.txt-active {
  background-color: #8155ba;
  color: white;
}

.btn:hover {
  color: black;
  background-color: #beafc2;
}

input[type=radio].btn-group {
  display: none;
}

input[type=radio].btn-group + label {
  color: white;
  padding: 1rem;
  /* margin: 1.5rem; */
  background-color: black;
  border: none;
  cursor: pointer;
  font-size: 1rem;
}

input[type=radio].btn-group.btn-active + label {
  background-color: #8155ba;
  color: white;
}

label {
  display: inline-block;
}

input[type=radio].btn-group.btn-first + label {
  border-radius: 0.25rem 0 0 0.25rem;
}

input[type=radio].btn-group.btn-last + label {
  border-radius: 0 0.25rem 0.25rem 0;
}

input[type=radio].btn-group + label:hover {
  color: black;
  background-color: #beafc2;
}
</style>
