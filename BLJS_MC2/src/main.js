const { Panel, Canvas, Button } = mc2;
const { Random, Color, Context } = bljs;

const width = 800;
const height = 800;
const demo = document.getElementById("demo");
const panel = new Panel(demo, 20, 20, 1140, 840);
const canvas = new Canvas(panel, 320, 20, width, height);

const context = canvas.context;
Context.extendContext(canvas.context);

render();

function render() {

}
