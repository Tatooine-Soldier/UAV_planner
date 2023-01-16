<template>
    <section id="calendar">
        <section class="calendar-container">
            <div class="month">
                <div><small @click="monthDown()">&#10094;</small> {{ months[currentDate.month] }} <small @click="monthUp()">&#10095;</small></div> <div><small>&#10094;</small> {{ currentDate.year }} <small>&#10095;</small></div>
            </div>
            <div class="weekdays">
                <div class="weekday" v-for="(weekday, index) in weekdays" :key="index">
                  {{ weekday }}
                </div>
            </div>
            <div class="dates">
                <div class="day-hidden" v-for="(n, index) in (firstMonthDay -1)" :key="'prev'+index">
                    {{ (prevMonthDays +1) - firstMonthDay + n }}
                </div>
                <div class="date" :class="{ active: n === currentDate.date}"
                v-for="(n, index) in currentMonthDays" 
                :key="'day'+index" @click="currentDate.date = n ;getBookings()" id="datecell"> {{ n }} </div>
                <div class="day-hidden" v-for="(n, index) in (43 - (currentMonthDays + firstMonthDay))" :key="'next'+index">
                    {{ n }}
                  </div>
            </div>
        </section>
    </section>
</template>

<style>
    .calendar-container {
        width: 99.4%;
        border: red solid 1px;
        background-color: rgb(100, 100, 100);

    }

    .month {
        padding: 5px;
    }

    .weekdays {
        display: grid;
        grid-template-columns: 14.28% 14.28% 14.28% 14.28% 14.28% 14.28% 14.28%;
    }

    .weekday {
        border: solid white 1px;
    }

    .month {
        display: grid;
        grid-template-columns: 50% 50%;
    }

    .date {
        cursor: pointer;
    }
    
    .date:hover {
          background-color: #fff;
          color: #2A4C50;
    }

    .date:nth-child(7n) {
          color: #D43541;
    }

    .dates {
        display: grid;
        grid-template-columns: 14.2% 14.2% 14.2% 14.2% 14.2% 14.2% 14.2%;
        grid-template-rows: 16% 16% 16% 16% 16% 16%;
    }

    .day-hidden {
        opacity: .4;
      }
    
    .day-hidden:nth-child(7n) {
          color: #D43541;
    }

      
    small:hover {
        cursor: pointer;
    }

</style>

<script>
    import axios from 'axios';
    export default {
      data: function() {
        return {
          weekdays: ['Mo','Tu','We','Th','Fr','Sa','Su'],
          months: [
            'January','February','March','April','May','June','July',
            'August','September','October','November','December'
          ],
          years: [
            '2023', '2024', '2025'
          ],
          currentDate: {
            date: 0,
            month: 0,
            year: 0
          },
          bookedDates: null
        }
      }, 
      props: ['propdates'],
      methods: {
        getCurrentDate() {
          let today = new Date();
          this.currentDate.date = today.getDate();
          this.currentDate.month = today.getMonth();
          console.log("MONTH", this.currentDate.month)
          this.currentDate.year = today.getFullYear();
        },
        monthUp() {
          if(this.currentDate.month === 11) {
            this.currentDate.month = 0;
            this.currentDate.year += 1;
          }
          else {
            this.currentDate.month += 1;
          }        
        },
        monthDown() {
          if(this.currentDate.month === 0) {
            this.currentDate.month = 11;
            this.currentDate.year -= 1;
          }
          else {
            this.currentDate.month -= 1;
          } 
        }, 
        getBookings() {
          // let date = document.getElementById("datecell");
          // var d = date.innerHTML
          // if (d.length != 2) {
          //   d = "0"+d
          // }

          let date = this.currentDate.date
          String(date)
          if (['1', '2', '3', '4', '5', '6', '7', '8', '9'].indexOf(date) >= 0) {
            date = "0"+date;
          }

          let month =  this.currentDate.month 
          month = month+1
          String(month)
          if (month.length != 2) {
            month = "0"+month
          }

          let year = this.currentDate.year

          let fullDate = year+"-"+month+"-"+date
          const queryDate = {
            date: fullDate,
          }
          console.log(queryDate)


        axios
          .post("/getDateFlight", queryDate)
          .then((response) => {
            const data = response.data;
            console.log("FETCHED FLIGHTS FOR THIS DATE: ",data);
          })
          .catch (function (error) {
              console.log("ERROR:", error);    
          })

        }
      },
      
      computed: {
        prevMonthDays() {
          let year = this.currentDate.month === 0 ? this.currentDate.year - 1 : this.currentDate.year;
          let month = this.currentDate.month === 0 ? 12 : this.currentDate.month;
          return new Date(year, month, 0).getDate();
        },
        firstMonthDay() {
          let firstDay = new Date(this.currentDate.year, this.currentDate.month, 1).getDay();
          if(firstDay === 0) firstDay = 7;
          return firstDay;
        },
        currentDay() {
          return new Date(this.currentDate.year, this.currentDate.month, this.currentDate.date).getDay();
        },
        currentMonthDays() {
          return new Date(this.currentDate.year, this.currentDate.month +1, 0).getDate();
        }
      }
      ,created() {
        this.getCurrentDate();
        // this.bookedDates = props.propdates;
      },
      // setup(props) {
      //   watch(() => props.propdates, () => {
      //     console.log("***PROPS***", props.propdates)
      //     var arrHours = ref([]);
      //     var arrDates = ref([]);  
      //     var date;
      //     for (date in props.propdates) {
      //       console.log(props.propdates[date])  //need to link these to the calendar. check console. maybe give each cell in the calendar a value and then match?
      //       var arrMy =  props.propdates[date].split(" ");
      //       arrDates.value.push(arrMy[1])
      //       arrHours.value.push(arrMy[0])
      //     }
      //     console.log("arrDates:", arrDates.value.length, arrHours.value.length)
      //     return {
      //       arrDates, arrHours
      //     }
      //   })
      // }
    }
</script>