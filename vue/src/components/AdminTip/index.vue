<template>
  <div id="explanation" :class="iocnClass" :style="style1" style="color:rgb(44, 188, 163); background-color: rgb(237, 251, 248)">
        <div id="checkZoom" class="title"><i class="fa fa-lightbulb-o"></i>
  				<h4 title="提示相关设置操作时应注意的要点" @click="zhanKai">操作提示</h4>
  				<span v-show="play" @click="clickCollapse" title="收起提示" id="explanationZoom" style="display: block;"></span>
  			</div>
  			<ul>
  				<slot></slot>
  			</ul>
  		</div>
</template>
<script>

export default {
  name: 'AdminTip',
  props: {
    operation: {
      type: String,
      default: ''
    },
    className: {
      type: String,
      default: ''
    }
  },
  computed: {
    iocnClass() {
      if (this.className) {
        return 'explanation ' + this.className
      } else {
        return 'explanation'
      }
    }
  },
  watch: {
    play: function(newAction) {
      var box = document.getElementById('explanation')
      var boxHeight = box.offsetHeight
      var boxWidth = box.offsetWidth
      if (newAction == false) {
        this.$TweenMax.fromTo(box, 0.3, { height: boxHeight, width: boxWidth }, { height: 31, width: 102 })
      } else {
        this.$TweenMax.to(box, 0.3, { width: '100%' })
        this.$TweenMax.to(box, 0.3, { height: '100%', delay: 0.3 })
      }
    }
  },
  methods: {
    clickCollapse: function(event) {
      this.play = !this.play
      console.log(this)
    },
    zhanKai: function() {
      if (this.play == true) return
      this.play = true
    }
  },
  data() {
    return {
      play: true,
      style1: { height: '100%', width: '100%' }
    }
  }
}
</script>

<style>
 *{word-wrap: break-word;
    outline: none;}
ul,li,h4{list-style-image: none;
    list-style-type: none;padding: 0;margin: 0;}
.explanation {font-size:12px;color: #0ba4da !important; background-color: rgba(79, 192, 232, 0.11) !important; display: block; padding: 6px 9px; border-radius: 5px; position: relative; overflow: hidden;}
.explanation:before{content: "";background-image: url(img/wave.png);width: 100%;height: 100%;position: absolute;top: 0px;left: 0px;border-radius: 5px;background-repeat: no-repeat;background-size: cover;}
.explanation .title {white-space: nowrap;margin-bottom: 8px;position: relative;cursor: pointer;}
.explanation .title h4 {font-size: 14px;font-weight: normal;line-height: 20px;height: 20px;display: inline-block;}
.explanation .title i { font-size: 18px; vertical-align: middle; margin-right: 6px; }
.explanation .title span {background-image: url(img/zhedie.png);width: 20px; height: 20px; position: absolute; z-index: 1; top: -6px; right: -9px; }
.explanation ul { color: #748A8F; margin-left: 10px; }
.explanation li { line-height: 20px; background: url(img/macro_arrow.gif) no-repeat 0 10px; padding-left: 10px; margin-bottom: 4px;  }
.fa-lightbulb-o:before {
    content: "";
    background: url(img/handd.png) no-repeat;
    width: 18px;
    height: 20px;
    display: inline-block;
}
</style>
